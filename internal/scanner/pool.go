package scanner

import (
	"fmt"
	"sync"

	"github.com/VirajD18/postgres-cis-scanner/internal/dashboard"
	"github.com/VirajD18/postgres-cis-scanner/internal/models"
)

func Run(servers []models.Server, workers int) {

	jobs := make(chan models.Server)

	var wg sync.WaitGroup
	var mu sync.Mutex

	var reports []dashboard.ServerReport

	for i := 1; i <= workers; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			for server := range jobs {

				report, err := Scan(server)

				if err != nil {

					fmt.Printf("[Worker %d] ERROR: %v\n", id, err)

					continue
				}

				mu.Lock()
				reports = append(reports, report)
				mu.Unlock()

			}

		}(i)

	}

	for _, server := range servers {
		jobs <- server
	}

	close(jobs)

	wg.Wait()

	err := dashboard.Generate(reports)

	if err != nil {
		fmt.Println(err)
	}

}
