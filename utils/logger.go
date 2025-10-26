package utils

// home grown logger, stores message so that gui can display
//
//type LogMsgStatus struct {
//	prefix string
//}
//
//var (
//	LogMsgStatusError = LogMsgStatus{"ERR"}
//	LogMsgStatusInfo  = LogMsgStatus{"INF"}
//)
//
//type LogMsg struct {
//	Status  LogMsgStatus    // ERR or INF
//	Message string          // String to display .. "" for Sink message.
//	Sink    chan<- []LogMsg // The chan to recieved all the messages, since last call.
//}
//
//var loggerChan chan LogMsg = make(chan LogMsg)
//
//func StartLogger() {
//
//	go func() {
//		var messages []LogMsg
//		loop := true
//		for loop {
//			msg := <-loggerChan
//			if msg.Message != "" {
//				messages = append(messages, msg)
//			} else if msg.Sink != nil {
//				togo := messages
//				messages = nil
//				msg.Sink <- togo
//			} else {
//				loop = false
//			}
//		}
//	}()
//
//}
//
//func StopLogger() {
//	loggerChan <- LogMsg{LogMsgStatusInfo, "", nil}
//}
//func LogInfo(msg string) {
//	loggerChan <- LogMsg{LogMsgStatusInfo, msg, nil}
//}
//
//func LogError(msg string) {
//	loggerChan <- LogMsg{LogMsgStatusError, msg, nil}
//}
//
