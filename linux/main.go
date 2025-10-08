package linux

/*

	- ubuntu DE er new update droto accept korena, fedora kore
	- ubuntu repo fedorar tulonai onk boro, ubutu te repo er baireo software install kora jai, jeta fedorate kora jaina


	- kde gnome teke km resource use kore , tobe eti user upor depend kore.
	- ubuntu debian er upor base kore toiri, fedora red hat er upor base kore toiri
		- ubuntu package manager/software manager = apt, fedora package manager/ software manager = dnf
		- software files format = .deb vs .rpm
	- ubuntu te বিশাল কমিউনিটি এবং অনলাইন টিউটোরিয়ালের সাপোর্ট পাওয়া
	- ফেডোরা DE er latast সংস্করণ খুব দ্রুত provide করে।
	- NVIDIA ড্রাইভার বা মিডিয়া কোডেক ইনস্টল করার জন্য fedora te প্রথমে RPM Fusion-এর মতো থার্ড-পার্টি রিপোজিটরি যোগ করতে হবে।
	-  ফেডোরা ডিফল্টভাবে বেশ হালকা (lean) থাকে। এটিতে Snap-এর মতো ভারী ব্যাকগ্রাউন্ড সার্ভিস ডিফল্টভাবে থাকে না। তারা Flatpak-কে বেশি গুরুত্ব দেয় যা তুলনামূলকভাবে হালকা।

	- ram usage difference of fedora vs ubuntu => 50-100 MB
		- ram usage difference of kde vs gnome => 200-400 MB
		- a browser consume ram => some GB
		- cpu usage difference of kde vs gnome => very small
		- battery usage difference of kde vs gnome => very small
			- kde te caile animation komie battery usage komano jai.








	- ল্যাপটপের ব্যাটারি লাইফের ক্ষেত্রে, Wayland ডিসপ্লে সার্ভার ব্যবহারে উভয় ডেস্কটপই X.Org-এর চেয়ে ভালো পাওয়ার এফিসিয়েন্সি দেখিয়েছে। (explain)
	- Ubuntu (Default): গ্নোম (GNOME)
	Fedora (Default): গ্নোম (GNOME)

	উভয় ডিস্ট্রিবিউশনেরই কেডিই প্লাজমাসহ অন্যান্য ডেস্কটপ এনভায়রনমেন্টের জন্য আলাদা সংস্করণ (উবুন্টুর ক্ষেত্রে "Flavours" এবং ফেডোরার ক্ষেত্রে "Spins") রয়েছে।  (explain)

	- windows = kernel + DE, linux = only kernel , right?
	- debian & red hat ki company?
	- দুটিই চমৎকার অপারেটিং সিস্টেম। আপনার প্রয়োজন অনুযায়ী যেকোনো একটি বেছে নিতে পারেন। সবচেয়ে ভালো হয়, ইনস্টল করার আগে দুটোকেই লাইভ ইউএসবি (Live USB) দিয়ে চালিয়ে দেখা। (explain)
	- সুতরাং, আপনি যদি কেডিই প্লাজমার সর্বশেষ সংস্করণ ব্যবহার করতে চান, তাহলে ফেডোরার কেডিই স্পিন (Fedora KDE Spin) আপনার জন্য অন্যতম সেরা একটি পছন্দ।
	আরেকটি বিকল্প: আপনি যদি একেবারে কেডিই ডেভেলপারদের তৈরি করা লেটেস্ট সংস্করণ ব্যবহার করতে চান, তাহলে কেডিই নিওন (KDE Neon) নামে একটি অপারেটিং সিস্টেম রয়েছে। এটি উবুন্টুর উপর ভিত্তি করে তৈরি, কিন্তু এর মূল উদ্দেশ্যই হলো ব্যবহারকারীদের কেডিই প্লাজমার একেবারে নতুন এবং স্থিতিশীল সংস্করণ সরবরাহ করা। (explain)







*/

/*



	ubuntu, fedora eder nijosso repository ache.
	jdi ubuntu repo te obs paoa na jai tahle :


	1) PPA hlo developer der repository.
	jkn tumi command die ppa k systemer sate add kro , tkn tumi asole apt er phobebook e obs er developerder repositoryr tikana add kore dao.
	systemer package manager (apt) sei tikana teke obs package ene install kore(defualt format = .deb format e)(jdi ppa te obs thake.)


	2) flatapk ekti univarsal package system (ortat eti ubuntu , fedora, arc liux soho sb linux distribution e cole.)
	developer ra tader apps flatpak hisebe "flathub" namok central store e release kore.
	tai ubuntu repo te jdi obs na thake tahle flathub teke "flatpak install..." command die install kora jai (jdi flatpak e obs thake)


	3) snap o flatapk er moto univarsal package system(eta ubuntur nirmata Canonical banaise, tai eta ubuntur sate vlovabe integreted)
	"snap store" teke "sudo snap install..." command die install kora jai


	### flathub, snap store google play store er mto cinta krte paro, jkane hajar hajar app thake (app package thake)
	developer ra অ্যাপগুলোকে তাদের প্রয়োজনীয় সবকিছুসহ(library, depency) একটিমাত্র প্যাকেজে দিয়ে দেয়। tai delete korar smy shudu app dlt krle hoina, pora package dlt krte hoi, dlt er smy command e package name dite hoi.
	flatapk , snap hlo play store er moto , sb developer ra ekane apps rake tai tumi sb rokomer apps pabe. ppa hlo obs develoerder repo, vlc developerder repo erkm. tai shudu obs er package add krle sekan teke only obs package pabe



	4) onk developer tader website e directly .deb file die rake . double click kore install kora jai or command "sudo dpkg -i obs-studio-version.deb"


	### apt die install krle app ti .deb package hisebe install hoi.
	snap ba flatpak die krle .snap ba .flatpak package (isolated/sandboxed) hisebe install hoi.

	* .deb er vetor only code thake , library & dependences system teke dhar nei.
	onnodike .snap ba .flatpak er vetorei library & dependencies thake. (ejonno isolated/sandboxed bole)
	ekon ekti mallicious .deb install hoye gele seta system er sb kicur access peye jabe . so secured na .
	ubuntu te .deb non-isolated ortat system access pai
	fedora te .rpm non-isolated ortat system access pai
	ar .snap & .flatpak holo modern & isolated jegulo sb distribution e cole.

*/
