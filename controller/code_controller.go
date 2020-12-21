package controller

import (
	"FoundationHelper_KnightHacks2020/model"
	// "bytes"
	"context"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// SyncBuf data structure
// type SyncBuf struct {
// 	mu       sync.Mutex
// 	buf      bytes.Buffer
// 	output   string
// 	overflow bool
// }

// Reset function resets the buffer
// func (s *SyncBuf) Reset() {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.buf.Reset()
// }

// func (s *SyncBuf) String() string {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	//log.Println("Buff Length:", s.buf.Len())
// 	if s.buf.Len() > 10000 {
// 		s.overflow = true
// 		s.buf.Truncate(10000)
// 	}
// 	return s.buf.String()
// }

// func (s *SyncBuf) Write(p []byte) (int, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	return s.buf.Write(p)
// }

func copyAndCapture(r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			//log.Println("N?", n)
			d := buf[:n]
			out = append(out, d...)
			if len(out) > 1024 {
				return out, err
			}
			// _, err := w.Write(d)
			// if err != nil {
			// 	return out, err
			// }
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

// RunCode function to get a bike
func RunCode(c *fiber.Ctx) error {
	// New incoming code
	userCode := new(model.UserCode)

	// Parse body into struct
	if bodyErr := c.BodyParser(userCode); bodyErr != nil {
		return c.Status(400).SendString(bodyErr.Error())
	}
	//log.Println(userCode.Input)

	f, fileErr := os.Create("./testproblems/linked_lists/user.c")
	if fileErr != nil {
		return c.Status(500).SendString(fileErr.Error())
	}
	_, writeErr := f.WriteString(userCode.Input)
	if writeErr != nil {
		return c.Status(500).SendString(writeErr.Error())
	}

	// Limit the bash script to only run for 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "runTest.sh")

	cmd.Dir = "testproblems/linked_lists/"

	var stdout []byte
	var errStdout error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	err := cmd.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(stdoutIn)
		wg.Done()
	}()
	var stderr []byte
	var errStderr error
	_ = stderr
	_ = errStderr
	stderr, errStderr = copyAndCapture(stderrIn)

	wg.Wait()

	err = cmd.Wait()

	if err != nil {
		log.Println("GOT ERROR:", err.Error())
		var errOutput string
		if err.Error() == "exit status 1" {
			errOutput = err.Error()
			if strings.Contains(errOutput, "Segmentation fault") {
				errOutput = "Segmentation Fault: You might have forgot to check for NULL"
			}
		} else {
			log.Println("Timed out?")
			errOutput = "Process Timed Out (Code exceeded 3 Seconds)"
		}
		return c.Status(400).JSON(fiber.Map{"error": errOutput})
	}

	result := fiber.Map{"output": stdout}
	return c.JSON(result)
}

// SimpleCode Function to just run some c proogram
func SimpleCode(c *fiber.Ctx) error {

	// New incoming code
	userCode := new(model.UserCode)

	// Parse body into struct
	if bodyErr := c.BodyParser(userCode); bodyErr != nil {
		return c.Status(400).SendString(bodyErr.Error())
	}

	f, fileErr := os.Create("./testproblems/testing.c")
	if fileErr != nil {
		return c.Status(500).SendString(fileErr.Error())
	}
	_, writeErr := f.WriteString(userCode.Input)
	if writeErr != nil {
		return c.Status(500).SendString(writeErr.Error())
	}

	// Limit the bash script to only run for 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "runC.sh")

	cmd.Dir = "testproblems/"

	var stdout []byte
	var errStdout error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	err := cmd.Start()
	log.Println("Started process")
	if err != nil {
		log.Println("Start error")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(stdoutIn)
		wg.Done()
	}()

	stderr, errStderr := copyAndCapture(stderrIn)

	wg.Wait()

	err = cmd.Wait()

	if stderr != nil || errStderr != nil {
		log.Println("Some stderr or errStderr")
		log.Println("Stderr:", stderr)
		log.Println("ErrStderr:", errStderr)
	}
	sendBack := string(stdout)
	if len(stdout) > 0 {
		log.Println("Output:", sendBack)
	}
	result := fiber.Map{"output": sendBack}
	return c.JSON(result)
}

// GetCode Function to get problem
func GetCode(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"problem": `#include <stdlib.h>
#include <stdio.h>

typedef struct node {
	int data;
	struct node* next, *prev;
} node;
		 
typedef struct queue {
	int size;
	struct node *head, *tail;
} queue;
		
int dequeue(queue* thisQ) {
	// Insert Code Here










}`,
		"problem_topic":   "Linked Lists",
		"problem_name":    "Fall 2020 Part A Question 2",
		"problem_summary": "Suppose we have a queue implemented as a doubly linked list using the structures shown below.  Use head for the front of the queue and tail for the end of the queue.\n\nstruct node {\n\tint data;\n\tstruct node* next, *prev;\n}\n\nstruct queue {\n\tint size;\n\tstruct node *head, *tail;\n}\n\nWrite a dequeue function for this queue. If the queue is NULL or is already empty, return 0 and take no other action. If the queue isn't empty, dequeue the appropriate value, make the necessary adjustments, and return the dequeued value. (Note: You must free the node that previously stored the dequeued value.)"})
}
