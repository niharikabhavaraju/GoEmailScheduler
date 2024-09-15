package scheduler

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/niharikabhavaraju/go_scheduler/pkg/models"
	"github.com/niharikabhavaraju/go_scheduler/pkg/utils"
)

var TickerDuration time.Duration

func init() {
	durationStr := os.Getenv("TICKER_DURATION")
	duration, _ := strconv.Atoi(durationStr)
	TickerDuration = time.Duration(duration) * time.Second
}

func StartScheduler() {
	// create a ticker
	ticker := time.NewTicker(TickerDuration)
	// create a go routine to send emails
	go func() {
		for {
			checkForActions()
			<-ticker.C
		}
	}()

}

func checkForActions() {
	t := time.Now()
	endTime := t.Format("2006-01-02 15:04:05")
	t1 := t.Add(-1 * TickerDuration)
	startTime := t1.Format("2006-01-02 15:04:05")
	log.Println("Checking for emails to be sent between", startTime, "and", endTime)
	emails, err := models.GetEmailByTime(startTime, endTime)
	if err != nil {
		log.Fatalf("Failed to get emails: %v", err)
	}
	if len(emails) > 0 {
		log.Println("Emails to be sent", emails)
		for _, email := range emails {
			err := utils.SendEmail("bhavarajuniharika@gmail.com", email.To, email.Subject, email.Body)
			if err == nil {
				log.Println("Email sent successfully")
				models.UpdateEmailStatus(&email, "sent")

			} else {
				log.Fatalf("Failed to send email: %v", err)
				models.UpdateEmailStatus(&email, "not sent")
			}
		}
	} else {
		log.Println("No emails to be sent")
	}

}
