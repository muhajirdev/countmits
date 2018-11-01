class Countmits < Formula
  desc ""
  homepage ""
  url "http://github.com/muhajirframe/countmits/releases/v0.2.2/countmits_0.2.2_Darwin_x86_64.tar.gz"
  version "0.2.2"
  sha256 "3525971474383427288fb28d1d01acc4e52512513143566cc205a0e33b4290c2"
  
  depends_on "git"
  depends_on "zsh"

  def install
    bin.install "countmits"
  end
end
