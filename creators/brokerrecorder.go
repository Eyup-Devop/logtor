// Package logtor provides log creators and loggers for various destinations.
//
// It includes implementations for creating logs and logging messages to a Kafka broker, a base log creator,
// and a central log manager (Logtor) that coordinates multiple log creators.
//
// The package leverages the "sarama" library for Kafka-related functionalities.
package creators

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/Eyup-Devop/logtor/types"
	"github.com/IBM/sarama"
)

// NewBrokerCreator creates a new instance of BrokerCreator, which logs messages to a Kafka broker.
//
// It initializes a BrokerCreator with the provided Kafka broker addresses, topic, time zone, log creator name, and call depth.
//
// Parameters:
//   - brokers: A list of Kafka broker addresses.
//   - topic: The Kafka topic to publish log messages.
//   - logName: The name representing the log creator (e.g., Broker).
//   - callDepth: The call depth to be used in log output.
//
// Returns:
//   - *BrokerCreator: A pointer to the newly created BrokerCreator.
//   - error: An error if initialization fails, or nil if successful.
func NewBrokerCreator(brokers []string, topic string, logName types.LogCreatorName, callDepth int) (*BrokerCreator, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write log entry:", err)
		}
	}()

	if logName == "" {
		logName = Broker
	}

	brokerCreator := &BrokerCreator{
		logName:   logName,
		topic:     topic,
		producer:  producer,
		callDepth: callDepth,
	}

	return brokerCreator, nil
}

// Broker is a constant representing the LogCreatorName for the Broker log creator.
const Broker types.LogCreatorName = "Broker"

// BrokerCreator is an implementation of the LogCreator interface for logging messages to a Kafka broker.
type BrokerCreator struct {
	producer  sarama.AsyncProducer
	topic     string
	logName   types.LogCreatorName
	callDepth int
}

// BrokerMessage represents the structure of log messages to be sent to the Kafka broker.
type BrokerMessage struct {
	LogLevel   string      `json:"loglevel"`
	Created    string      `json:"created"`
	File       string      `json:"file"`
	Line       int         `json:"line"`
	LogMessage interface{} `json:"log_message"`
}

// LogItWithCallDepth logs a message with the specified log level, call depth, and log message to the Kafka broker.
//
// It formats the log entry with the log level, timestamp, file name, line number, and log message,
// then sends the formatted JSON message to the Kafka broker.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - callDepth: The call depth for recording the log entry.
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (br *BrokerCreator) LogItWithCallDepth(level types.LogLevel, callDepth int, logMessage interface{}) bool {
	var (
		file string
		line int
		ok   bool
	)
	_, file, line, ok = runtime.Caller(callDepth)
	if !ok {
		file = "UNKNOWN FILE"
		line = 0
	}

	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006/01/02 15:04:05")

	message := BrokerMessage{
		LogLevel:   string(level),
		Created:    formattedTime,
		File:       file,
		Line:       line,
		LogMessage: logMessage,
	}

	jsonMessage, _ := json.Marshal(message)

	br.producer.Input() <- &sarama.ProducerMessage{
		Topic: br.topic,
		Key:   sarama.StringEncoder("0"),
		Value: sarama.ByteEncoder(jsonMessage),
	}
	return true
}

// LogIt logs a message with the specified log level using the default call depth to the Kafka broker.
//
// This method is a convenience wrapper around LogItWithCallDepth, using the call depth configured for the BrokerCreator instance.
//
// Parameters:
//   - level: The log level for the message (e.g., INFO, DEBUG).
//   - logMessage: The message to be logged, which can be of any type.
//
// Returns:
//   - bool: Always returns true, indicating the message was successfully logged.
func (br *BrokerCreator) LogIt(level types.LogLevel, logMessage interface{}) bool {
	return br.LogItWithCallDepth(level, br.callDepth, logMessage)
}

// LogName returns the name of the log creator.
//
// Returns:
//   - LogCreatorName: The name of the log creator.
func (br *BrokerCreator) LogName() types.LogCreatorName {
	return br.logName
}

// SetCallDepth sets the call depth for recording log entries.
//
// This method allows configuring how deep into the call stack the logger should trace when recording
// log messages. A higher call depth includes more layers of function calls in the log output,
// providing additional context about the log origin.
//
// Parameters:
//   - callDepth: The depth to set for recording log entries.
func (br *BrokerCreator) SetCallDepth(callDepth int) {
	br.callDepth = callDepth
}

// CallDepth returns the current call depth setting for recording log entries.
//
// Returns:
//   - int: The current call depth setting for recording log entries.
func (br *BrokerCreator) CallDepth() int {
	return br.callDepth
}

// Shutdown gracefully shuts down the BrokerCreator by closing the Kafka producer.
//
// Use this method to perform any necessary cleanup or shutdown operations for the log creator.
func (br *BrokerCreator) Shutdown() {
	br.producer.Close()
}
