package commend

import "gin-lib/app/service"

// DateNewlyCreate exec for cron and insert database
func DateNewlyCreate() {
	service.DateNewlyAddCreate()
}
