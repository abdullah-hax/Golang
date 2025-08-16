/*
			PC / Operating System                   | Go Process / Go Runtime
		----------------------------------------                    |----------------------------------------
		যখন PC চালু হয়, তখন প্রথমে OS     | যখন একটি Go process (virtual pc) চালু হয়, তখন প্রথমে
	    চালু হয় এবং পরে পুরো PC-র control   | Go runtime (virtual OS) চালু হয়
		নিয়ে নেয়।                         | এবং পরে process-এর control করে
		OS kernel-এর 'scheduler' thread-কে         | Go runtime-এর 'goroutine scheduler'
		(execute) করে।                       | goroutine(virtual thread)-কে execute করে।
		PC-তে যতগুলো virtual CPU (vCPU) থাকে,    | Go runtime-এও ঠিক ততগুলো vCPU create হয়
												| এবং mapping হয়।


		goroutine বেড়ে গেলে scheduler go process এর ভেতর আরও vcpu ক্রিয়েট করে এবং একই সাথে pc তেও vcpu ক্রিয়েট হয় এবং mapping হয়.

	    "এটি সম্পূর্ণ ১০০% নিখুঁত নয়, তবে শিক্ষামূলক উদ্দেশ্যে এভাবে আমি PC-এর OS এবং Go runtime-এর মধ্যে মিলটি বুঝি।"

		keyword :
		## process = virtual computer (প্রত্যেকটা process এর ভেতর থাকে virtual porcessor)
		thread(যখন একটা থাকে) = virtual process
		go runtime = virtual OS (mini OS)
		go routine = virtual thread



		# os calo hoar pr onk kisu initialize kore
		tmni go runtime o calo hoar pr onk kisu initialize kore
		1. Go routine scheduler
		2. Heap allocator
		3. garbage collector (It deletes heap's ajaira data)
		4. logical processor


		# go routine create kore programmar ra
		# thread er stack 8MB, goroutine er stack 2KB


		NB : really code execute kore cpu, os noi
		     cpu -> os -> os er kernel er scheduler thread -> go runtime/code
			 jdio cpu soho sbkisu os er control e thake kintu os k muloto cpu i run kore(karon os nijei ekta code, ar code execute cpu chara keo krte parena).
*/
package main

import (
	"fmt"
	"time"
)

var a = 30

const b = 40

func stack(num int) {
	fmt.Println("num = ", num)
}

func main() {
	var c int = 10
	fmt.Println("hello", c)

	go stack(1)
	go stack(2)
	go stack(3)
	go stack(4)
	go stack(5)

	fmt.Println(a, " ", b)

	time.Sleep(5 * time.Second)

	fmt.Println("ha ha")

}

// total goroutine 6ta = 6ta stack
// 1ta main goroutime + 5ta goutine = 6 ta goroutine
