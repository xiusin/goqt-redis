#!/usr/local/bin/v run
defer {
	rm("build")
}
println ('编译中...')

exec("rm -rf deploy")?

exec("qtdeploy build desktop") or {
	panic(err)
}

println ('执行...')

mkdir("deploy/darwin/goqt-redis.app/Contents/MacOS/dist")?

cp_all("dist", "deploy/darwin/goqt-redis.app/Contents/MacOS/dist", true) or {
    panic(err)
}

exec("qtdeploy run desktop") or {
	panic(err)
}

