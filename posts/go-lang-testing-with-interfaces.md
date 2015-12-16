# Go Lang - Testing With Interfaces

You cannot go far with any programming language or framework without writing some tests. Sooner or later your application will grow, so that you have to have some kind of help to keep it working correctly. When I started working with Go language it was one of the first things I wanted to try. So ho exactly do we write tests in Go?

### Testing patterns

There are two main approaches to testing a piece of software. You can have some function that takes some input data structure (often primitives, but also some kind of maps/arrays, data objects), perform some kind of calculations on it and return some checkable result. The second approach is you create what I call computation point - a place where you inject your logic-performing functions and see not only for the result iself, but also for correct functions calls. Let's take a closer look on both of them.

Before we start, let's create a foundation for our tests in a new file. Imagine that we have a file called `gradecalculator.go`, which is responsible for calculating thei final grade for a student based on their homework grades:

    package calculator
    
    func FinalGrade(grades []int) int {
    	return 5
    }
    
Now this looks silly, but let's say that we are the best teacher ever and we give everybody a grade of five (in Poland it's an equivalent for A in US). Just for the record, we want to make sure that this stays unchanged after any changes we might make in the future, so we want to test it. All we need to do is create a file called `gradecalculator_test.go`, import a module called `testing` and...

### Simple input-output tests

The first approach is often refered to as **blackbox** testing. You don't care what is happening inside your function - the most important thing is the final result. So if we want to test out `FinalGrade` function, all we need to do is write:

    package calculator_test
    
    import (
    	"testing"
    
    	"github.com/mycodesmells/go-tutorial/calculator"
    )
    
    func TestFinalGradeShouldEqualToFive(t *testing.T) {
    	fg := calculator.FinalGrade([]int{5})
    
    	if(fg != 5) {
    		t.Errorf("Calculated incorrect final grade. Expected 5 but was %v", fg)
    	}
    }

You have to remember a couple of things here. First, you absolutely need to import `testing` package and use `*testing.T` as a parameter for your tests. This is necessary for printing any test errors that might occur. Second thing is to have your test functions begin with `Test*`, so that they are actually getting called by testing command. Last thing is optional in theory, but it's a pretty good practice to have your test in a `xyz_test` package when testing stuff from `xyz` package. This allows your test to access only the public functions and objects from the production files, so that you don't have any higher privileges in your tests than other files in your project.

Running this test results in:

    $ go test ./...
    ?       github.com/mycodesmells/go-tutorial     [no test files]
    ok      github.com/mycodesmells/go-tutorial/calculator  0.001s
    ?       github.com/mycodesmells/go-tutorial/database    [no test files]

As you can see, we only have tests in `go-tutorial/calculator` directory (shame on us!), but at least those tests pass.

### Mocking stuff

The second approach we need to have in our arsenal is testing the behaviour of some passed-by objects. This might sound hard, but in fact it's not at all difficult. All you need to do is have an interface as your parameter instead of a struct. Imagine we want to save our final grades in the database, so we want to have a teacher calculating the grade and the result will be stored in DB. Our desired behaviour would look like this:

    func SaveFinalGrade(t people.Teacher, s people.Student) string
    
Our function will return a query we would be making to our fictional database. Now, let's create two interfaces for our parameters:

**people/teacher.go**

    type Teacher interface {
    	FinalGrade(grades []int) int
    }

**people/student.go**
    
    type Student interface {
    	Grades() []int
    	FullName() string
    }
    
And then we can create a test (#TDD) with our expected results:

    func TestShouldCreateFinalGradeQuery(t *testing.T) {
    	ft := FakeTeacher{}
    	fs := FakeStudent{}
    
    	query := database.SaveFinalGrade(ft, fs)
    	exp := "Giving 1 for Jerry Lemon"
    	if query != exp {
    		t.Errorf("Incorrect query. Expected %v but got %v", exp, query)
    	}
    }
    
To make if work, we need to create those `FakeTeacher` and `FakeStudent` objects. In Go, to implement an interface all you need to do is create a struct that has all the functions defined in it (_If it walks like a duck and quacks like a duck - it is a duck_ rule). Now, we don't want to test the logic of `people.Teacher` nor `people.Student`, but we want to know if the query would be built correctly. Assumption here is that both input objects do their work properly (with their own tests in separate place). So there they are:
    
    type FakeTeacher struct{}
    func (ft FakeTeacher) FinalGrade(grades []int) int {
    	return 1
    }
    
    type FakeStudent struct{}
    func (fs FakeStudent) FullName() string {
    	return "Jerry Lemon"
    }
    func (fs FakeStudent) Grades() []int {
    	return []int {5}
    }
    
Let's make sure that the test is failing:

    $ go test ./...
    ...
    --- FAIL: TestShouldCreateFinalGradeQuery (0.00s)
            database_test.go:33: Incorrect query. Expected Giving 1 for Jerry Lemon but got 
    FAIL
    FAIL    github.com/mycodesmells/go-tutorial/database    0.001s
    ...
    
Good. As you can see, we have one harsh teacher and a pretty good student. But out teacher implementation says clearly: no matter how good the grades are, we are giving `1` as the final grade. Now we can focus on changing the logic of our production code:

    func SaveFinalGrade(t people.Teacher, s people.Student) string {
    	gr := t.FinalGrade(s.Grades())
    	fn := s.FullName()
    	return fmt.Sprintf("Giving %d for %s", gr, fn)
    }
    
This time our tests run smoothly:

    $ go test ./...
    ?       github.com/mycodesmells/go-tutorial     [no test files]
    ok      github.com/mycodesmells/go-tutorial/calculator  0.001s
    ok      github.com/mycodesmells/go-tutorial/database    0.001s
    ?       github.com/mycodesmells/go-tutorial/people      [no test files]

Source code for this example is available [on GitHub](https://github.com/mycodesmells/go-tutorial). Note that you should clone repository into $GOPATH/src/github.com/mycodesmells/go-tutorial directory.
