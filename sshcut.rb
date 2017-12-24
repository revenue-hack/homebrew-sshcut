require "formula"

class Sshcut < Formula
  homepage "https://github.com/revenue-hack/homebrew-sshcut"
  url "https://github.com/revenue-hack/homebrew-sshcut.git"
  sha256 "606d00bc4736ef3fe10fdaa994985a08fd5642279a96f30ab2727b2cc7a771c1"
  head "https://github.com/revenue-hack/homebrew-sshcut.git"
  version "1.0.0"
  depends_on 'go' => :build

  def install
    ENV['GOPATH'] = buildpath
    #system 'go', 'get', '-u', 'github.com/golang/dep/cmd/dep'
    system 'go', 'get', 'github.com/revenue-hack/homebrew-sshcut'
    mkdir_p buildpath/'src/github.com/tmp'
    ln_s buildpath, buildpath/'src/github.com/tmp/homebrew-sshcut'
    system 'cd', buildpath/'src/github.com/tmp/homebrew-sshcut'
    system 'ls', '-la'
    #system 'dep', 'ensure'
    system 'go', 'get', 'github.com/mitchellh/go-homedir'
    system 'go', 'build', '-o', 'sshcut', buildpath/'src/github.com/tmp/homebrew-sshcut/main.go'
    bin.install 'sshcut'
  end
end
