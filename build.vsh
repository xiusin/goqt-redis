#!/usr/local/bin/v run

import term

defer {
	rm("build")
}

term.clear()

println ('编译中...')

res := exec("qtdeploy build desktop") or {
	panic(err)
}
if res.exit_code != 0 {
    eprintln(res.output)
    exit(1)
}

println ('执行...')

mkdir("deploy/darwin/goqt-redis.app/Contents/MacOS/dist")?

cp_all("dist", "deploy/darwin/goqt-redis.app/Contents/MacOS/dist", true) or {
    panic(err)
}
mkdir("deploy/darwin/goqt-redis.app/Contents/MacOS/qss")?

cp_all("qss", "deploy/darwin/goqt-redis.app/Contents/MacOS/qss", true) or {
    println("qss: ${term.fail_message(err)}")
    return
}

cp("goqt-redis.icns", "deploy/darwin/goqt-redis.app/Contents/Resources/goqt-redis.icns") or {
    println("icns: ${term.fail_message(err)}")
    return
}

exec("qtdeploy run desktop") ?

