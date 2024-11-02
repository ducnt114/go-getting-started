package demo_error

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var Cmd = &cobra.Command{
	Use:   "demo_error",
	Short: "demo_error",
	Long:  `demo_error`,
	Run: func(cmd *cobra.Command, args []string) {
		//demoErrorIs()
		//demoErrorAs()
		//demoCallerFrame()
		demoCallFrameWithCustomError()
	},
}

type StackError struct {
	Message string
	Frames  []uintptr
}

func (e *StackError) Error() string {
	// Use runtime.CallersFrames to format the stack trace
	frames := runtime.CallersFrames(e.Frames)
	stackTrace := ""
	for {
		frame, more := frames.Next()
		stackTrace += fmt.Sprintf("%s\n\t%s:%d\n",
			frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return fmt.Sprintf("%s\nStack trace:\n%s", e.Message, stackTrace)
}

func newStackError(msg string) error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	return &StackError{Message: msg, Frames: pc}
}

func doSomething() error {
	return newStackError("Something went wrong")
}

func demoCallFrameWithCustomError() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}

/*

 */

func demoCallerFrame() {
	printStack()
}

func printStack() {
	// Create a slice to store the program counters
	pc := make([]uintptr, 10)
	// Skip 2 frames: runtime.Callers and printStack itself
	n := runtime.Callers(2, pc)
	// Retrieve and format the call stack frames
	frames := runtime.CallersFrames(pc[:n])
	// Iterate over the frames and print information
	for {
		frame, more := frames.Next()
		fmt.Printf("Function: %s\nFile: %s\nLine: %d\n\n",
			frame.Function, frame.File, frame.Line)
		// Check if there are more frames
		if !more {
			break
		}
	}
}

/*****
error as
*/

func demoErrorAs() {
	err := operation()
	var myErr *MyError
	if errors.As(err, &myErr) { // Check if err is of type *MyError
		fmt.Printf("Error occurred: %s (Code: %d)\n", myErr.Message, myErr.Code)
	} else {
		fmt.Println("Unknown error")
	}
}

// MyError Custom error type
type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func operation() error {
	return &MyError{Message: "something went wrong", Code: 404} // Return custom error
}

// ***** error is *****

var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
	if id == 0 {
		return fmt.Errorf("lookup error: %w", ErrNotFound) // Wrap the error
	}
	return nil
}

func demoErrorIs() {
	err := findItem(0)
	if errors.Is(err, ErrNotFound) { // Check if the error is or wraps ErrNotFound
		fmt.Println("Error: Item not found")
	} else {
		fmt.Println("Item found")
	}
}
