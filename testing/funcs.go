package testing

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Size represents a width and height pair
type Size struct {
	Width  int
	Height int
}

//trace writes a message to the log if the
//TRACE environment variable is set to a
//true value
func trace(message string) {
	traceSetting := strings.ToLower(os.Getenv("TRACE"))
	if traceSetting == "1" || traceSetting == "true" {
		log.Println(message)
	}
}

//GetGreeting returns a greeting for a `name`.
//if name is zero-length, or all spaces,
//"World" will be used instead
func GetGreeting(name string) string {
	if len(name) == 0 {
		name := "World"
		trace(fmt.Sprintf("defaulting name to `%s`", name))
	}
	return fmt.Sprintf("Hello, %s!", name)
}

//ParseSize parses a string in the form of
//"WidthxHeight", returning a populated
//Size struct. If the Width or Height are
//not valid numbers, their respective fields
//will be set to zero
func ParseSize(size string) *Size {
	nums := strings.Split(size, "x")
	width, _ := strconv.Atoi(nums[0])
	height, _ := strconv.Atoi(nums[1])
	return &Size{width, height}
}

//DefaultLateDays is the default number of late days each netID gets
const DefaultLateDays = 4

//LateDays tracks the number of late days left for a given netID
type LateDays struct {
	entries map[string]int
}

//NewLateDays constructs a new LateDays tracker
func NewLateDays() *LateDays {
	return &LateDays{
		entries: map[string]int{},
	}
}

//Consume consumes a late day for a given netID
//and returns the number of late days left for
//that netID. If the netID has already consumed
//all late days, it continues to return 0.
func (ld *LateDays) Consume(netID string) int {
	//if netID is not yet a key in the map
	//add an entry with the value 4
	if ld.entries[netID] == 0 {
		ld.entries[netID] = DefaultLateDays
	}
	ld.entries[netID]--
	return ld.entries[netID]
}
