/*# Bob

Bob is a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying
anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it comes to written communication and 
always follows normal rules regarding sentence punctuation in English.*/
package bob

import "strings"
func Hey(remark string) string {

switch {

case question(remark): 
	return "sure"

case YELL(remark):
	return "Whoa, chill out!"

case question_yell(remark):
	return "Calm down, I know what I'm doing!"
	
case addressing_silent(remark):
	return "Fine. Be that way!"

default:
	return "Whatever."




	return ""
}
