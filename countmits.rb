class Countmits < Formula
  desc ""
  homepage ""
  url "http://github.com/muhajirframe/countmits/releases/download/v0.2.4/countmits_0.2.4_Darwin_x86_64.tar.gz"
  version "0.2.4"
  sha256 "9271abc636f301d04b2fe9a7aee5541e39bc5846cdfaec1a9ec6e644338f74ce"
  
  depends_on "git"

  def install
    bin.install "countmits"
  end
end
