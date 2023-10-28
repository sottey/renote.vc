desc "Builds renotevc for release"

Envs = [
  { goos: "darwin", arch: "386" },
  { goos: "darwin", arch: "amd64" },
  { goos: "linux", arch: "arm" },
  { goos: "linux", arch: "arm64" },
  { goos: "linux", arch: "386" },
  { goos: "linux", arch: "amd64" },
  { goos: "windows", arch: "386" },
  { goos: "windows", arch: "amd64" }
].freeze

Version = "0.5.0".freeze

task :build do
  `rm -rf dist/#{Version}`
  Envs.each do |env|
    ENV["GOOS"] = env[:goos]
    ENV["GOARCH"] = env[:arch]
    puts "Building #{env[:goos]} #{env[:arch]}"
    `GOOS=#{env[:goos]} GOARCH=#{env[:arch]} CGO_ENABLED=0 go build -v -o dist/#{Version}/renotevc`
    if env[:goos] == "windows"
      puts "Creating windows executable"
      `mv dist/#{Version}/renotevc dist/#{Version}/renotevc.exe`
      `cd dist/#{Version} && zip renotevc_win.zip renotevc.exe`
      puts "Removing windows executable"
      `rm -rf dist/#{Version}/renotevc.exe`
    else
      puts "Tarring #{env[:goos]} #{env[:arch]}"
      `cd dist/#{Version} && tar -czvf renotevc_#{env[:goos]}_#{env[:arch]}.tar.gz renotevc`
      puts "Removing dist/#{Version}/renotevc"
      `rm -rf dist/#{Version}/renotevc`
    end
  end
end

task default: :build
