package main

/* 


	11 এর binary = 1011

	32-bit PC-এর register cell-এ 11 একটি cell-এ থাকে, অর্থাৎ 32 digit আকারে থাকে।
	64-bit PC-এর register cell-এ 11 একটি cell-এ থাকে, অর্থাৎ 64 digit আকারে থাকে।
	তবে bit manipulation করলে একটি cell-এ অনেকগুলো ছোট সংখ্যা রাখা যায়।

	Architecture যতই হোক, RAM-এর cell always 8-bit।
	11 সংখ্যা RAM-এ কত digit আকারে থাকবে সেটা architecture-এর উপর নির্ভর করে না, বরং data type-এর উপর নির্ভর করে।
	যদি 11 int হয়, তবে এটি RAM-এ 32 digit আকারে থাকবে (4 byte)।

	*binary 1 digit = 1 bit

*/
