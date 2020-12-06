import java.io.File

var counter = 0
var counter_2 = 0

var alphabet = "qwertyuioopasdfghjklzxcvbnm"
var answeredQuestions = setOf<Char>()
var allAnsweredQuestions = alphabet.toSet()

File("input.txt").forEachLine {
	if (it == "") {
		counter += answeredQuestions.count()
		counter_2 += allAnsweredQuestions.count()
		answeredQuestions = setOf<Char>()
		allAnsweredQuestions = alphabet.toSet()
	} else {
		var currentSet = it.toSet()
		answeredQuestions += currentSet
		allAnsweredQuestions = allAnsweredQuestions intersect currentSet
	}
}

println(counter + answeredQuestions.count())
println(counter_2 + allAnsweredQuestions.count())
