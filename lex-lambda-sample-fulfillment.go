package main

import (
   "context"
   "encoding/json"
   "fmt"
   "github.com/aws/aws-lambda-go/events"
   "github.com/aws/aws-lambda-go/lambda"
   "go.uber.org/zap"
   "time"
)

type Response struct {
   Answer string `json:"Answer" yaml:"Answer"`
}

//"dialogAction": {
//   "type": "Close",
//   "fulfillmentState": "Fulfilled or Failed",
//   "message": {
//   "contentType": "PlainText or SSML or CustomPayload",
//   "content": "Message to convey to the user. For example, Thanks, your pizza has been ordered."
//},


func init (){
   masterLogger, err := zap.NewDevelopment()
   if err == nil {
      zap.ReplaceGlobals(masterLogger)
   }
}

//version variable which can be overidden at compile time
var (
   version = "dev-local-version"
   commit  = "none"
   date    = "unknown"
)

func HandleRequest(ctx context.Context, lexEvent events.LexEvent) (events.LexEvent, error) {
   var logger = zap.S()

   logger.Infof("Version git Tag: %s Date: %s Commit: %s", version, date, commit)

   jLexEvent, _ := json.Marshal(lexEvent)
   logger.Info(string(jLexEvent))


   loc,_ := time.LoadLocation("Australia/Sydney")


   event := events.LexEvent{
      DialogAction: &events.LexDialogAction{
         Type: "Close",
         FulfillmentState: "Fulfilled",
      },
   }

   event.DialogAction.Message = make(map[string]string,2)
   event.DialogAction.Message["contentType"] = "SSML"

   if lexEvent.CurrentIntent.Name == "DayNow" {
      event.DialogAction.Message["content"] = time.Now().In(loc).Weekday().String()
   } else {

      //timeStr := fmt.Sprintf("<speak><prosody rate=\"slow\"><amazon:auto-breaths>The time is %s.</amazon:auto-breaths></prosody></speak>", time.Now().In(loc).Format(time.Kitchen))
      event.DialogAction.Message["content"] = fmt.Sprintf("<speak><prosody rate=\"slow\"><amazon:auto-breaths>The time is %s.</amazon:auto-breaths></prosody></speak>", time.Now().In(loc).Format(time.Kitchen))
   }

   jLexEvent, _ = json.Marshal(event)
   logger.Info(string(jLexEvent))

   return event, nil

}

func main() {
   lambda.Start(HandleRequest)
}
