package cronjob

import (
	"fmt"
	"gblog/app"
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

type BaseCronJobHandler struct {
	App *app.Application
}

type CronJobHandler struct {
	Name      string
	Job       func()
	Schedule  string
	CronEntry cron.EntryID // Cron job 的唯一标识符
}

type CronJobManager struct {
	App             *app.Application
	CronJobHandlers []CronJobHandler
	CronJobEntries  map[string]cron.EntryID
	mu              sync.Mutex
	cronScheduler   *cron.Cron
}

func NewCronJobManager(app *app.Application) *CronJobManager {
	return &CronJobManager{
		App:            app,
		CronJobEntries: make(map[string]cron.EntryID),
		cronScheduler:  cron.New(),
	}
}

func (m *CronJobManager) RegisterCronJob(name, schedule string, job func()) {
	m.mu.Lock()
	defer m.mu.Unlock()

	entryID, err := m.cronScheduler.AddFunc(schedule, func() {
		log.Printf("Running cron job: %s", name)
		job()
	})

	if err != nil {
		log.Printf("Failed to add cron job %s: %v", name, err)
		return
	}

	m.CronJobHandlers = append(m.CronJobHandlers, CronJobHandler{
		Name:      name,
		Schedule:  schedule,
		CronEntry: entryID,
	})
	m.CronJobEntries[name] = entryID
	log.Printf("Cron job %s registered successfully.", name)
}

func (m *CronJobManager) Start() {
	m.cronScheduler.Start()
	log.Println("Cron jobs started.")
}

func (m *CronJobManager) GetCronJobList() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	var jobNames []string
	for _, handler := range m.CronJobHandlers {
		jobNames = append(jobNames, handler.Name)
	}
	return jobNames
}

func (m *CronJobManager) PauseCronJob(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	entryID, exists := m.CronJobEntries[name]
	if !exists {
		log.Printf("Cron job %s not found", name)
		return fmt.Errorf("cron job %s not found", name)
	}

	m.cronScheduler.Remove(entryID)
	log.Printf("Cron job %s paused.", name)
	return nil
}
func (m *CronJobManager) ResumeCronJob(name, schedule string, job func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	entryID, err := m.cronScheduler.AddFunc(schedule, func() {
		log.Printf("Running cron job: %s", name)
		job()
	})

	if err != nil {
		return err
	}

	m.CronJobEntries[name] = entryID
	log.Printf("Cron job %s resumed.", name)
	return nil
}

func (m *CronJobManager) Stop() {
	m.cronScheduler.Stop()
	log.Println("Cron jobs stopped.")
}

func (m *CronJobManager) StopCronJob(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	entryID, exists := m.CronJobEntries[name]
	if !exists {
		log.Printf("Cron job %s not found", name)
		return fmt.Errorf("cron job %s not found", name)
	}

	m.cronScheduler.Remove(entryID)
	delete(m.CronJobEntries, name)
	log.Printf("Cron job %s stopped.", name)
	return nil
}
