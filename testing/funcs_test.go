package testing

import (
    "testing"
)

func TestReverse(t *testing.T) {
    cases := []struct{
        name string
        input string
        expectedOutput string

    }{
        {
            name: "empty string",
            input: "",
            expectedOutput: "",
        },
        {
            name: "single char",
            input: "a",
            expectedOutput: "a",
        },
        {
            name: "two char",
            input: "ab",
            expectedOutput: "ba",
        },
        {
            name: "three char",
            input: "abc",
            expectedOutput: "cba",
        },
        {
            name: "palindrome?",
            input: "stressed",
            expectedOutput: "desserts",
        },
        {
            name: "palindrome",
            input: "racecar",
            expectedOutput: "racecar",
        },
        {
            name: "high unicode",
            input: "Hello, 世界",
            expectedOutput: "界世 ,olleH",
        },
        {
            name: "high unicode",
            input: "Hello 😋",
            expectedOutput: "😋 olleH",
        },
    }

    for _, c := range cases {
        if output := Reverse(c.input); output != c.expectedOutput {
            t.Errorf("%s: Got %s but expected %s", c.name, output, c.expectedOutput)
        }
    }
}

func TestGetGreeting(t *testing.T) {
    cases := []struct{
        name string
        input string
        expectedOutput string

    }{
        {
            name: "No input",
            input: "",
            expectedOutput: "Hello, World!",
        },
        {
            name: "Some input",
            input: "😋",
            expectedOutput: "Hello, 😋!",
        },
    }

    for _, c := range cases {
        if output := GetGreeting(c.input); output != c.expectedOutput {
            t.Errorf("%s: Got %s but expected %s", c.name, output, c.expectedOutput)
        }
    }
}

func TestParseSize(t *testing.T) {

}

func TestLateDaysConsume(t *testing.T) {
    ld := NewLateDays()
    for i := 3; i > -10; i-- {
        expectedOutput := i
        if expectedOutput < 0 {
            expectedOutput = 0
        }
        if output := ld.Consume("test"); output != expectedOutput {
            t.Errorf("iteration %d: got %d but expected %d", i, output, expectedOutput)
        }
    }
}
