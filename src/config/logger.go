package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// SetupLogger will configure both `logrus` and `lumberjack` and allow for log rotation and log file compression
func SetupLogger(logFilePath string) {
	// Get ENV Variables to configure how we handle log files
	maxFileSize, _ := strconv.ParseInt(os.Getenv("MAX_LOG_FILE_SIZE"), 10, 64)
	if maxFileSize == 0 {
		maxFileSize = 5
	}
	maxBackUps, _ := strconv.ParseInt(os.Getenv("MAX_FILE_BACKUPS"), 10, 64)
	if maxBackUps == 0 {
		maxBackUps = 10
	}
	maxAge, _ := strconv.ParseInt(os.Getenv("MAX_FILE_AGE"), 10, 64)
	if maxAge == 0 {
		maxAge = 30
	}
	compress, _ := strconv.ParseBool(os.Getenv("COMPRESS_TO_ZIP"))

	lumberjackLogger := &lumberjack.Logger{
		// Log file absolute path, os agnostic
		Filename:   filepath.ToSlash(logFilePath + "/" + GetLogFilename()),
		MaxSize:    int(maxFileSize), // in MBs
		MaxBackups: int(maxBackUps),
		MaxAge:     int(maxAge), // days
		Compress:   compress,    // disabled by default
	}

	// Fork writing into two outputs
	multiWriter := io.MultiWriter(os.Stderr, lumberjackLogger)

	logFormatter := new(log.JSONFormatter)       // you can also use log.JSONFormatter{} to allow log analysis
	logFormatter.TimestampFormat = time.RFC1123Z // writes in this format: "Mon, 02 Jan 2006 15:04:05 -0700"
	logFormatter.PrettyPrint = true

	log.SetFormatter(logFormatter)
	log.SetOutput(multiWriter)
}

// GetLogFilename creates the log file name with current date
func GetLogFilename() string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return "application-" + t.Format(layout) + ".log"
}

// CreateLogPaths will receive the log path variable $LOG_FILE_PATH and create the folder path
func CreateLogPaths() (string, error) {
	// String that contains the log path
	var logFilePath string = os.Getenv("LOG_FILE_PATH")

	if logFilePath == "" {
		log.Warnf("No ENV Variable set for $LOG_FILE_PATH is set, defaulting to /tmp/daraja path for logs")
		logFilePath = "/tmp/daraja"
	}

	err := os.MkdirAll(logFilePath, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to create logs path: %s due to error: %+v", logFilePath, err)
		return "", err
	}
	return logFilePath, nil
}
