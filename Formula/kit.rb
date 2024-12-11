# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kit < Formula
  desc "Crazy fast local dev loop."
  homepage "https://github.com/kitproj/kit"
  version "0.1.19"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/kitproj/kit/releases/download/v0.1.19/kit_0.1.19_Darwin_arm64.tar.gz"
      sha256 "c77ae06f0d2cdb8876933a86e6853f69972459cb45d38d88ec431a430c2b5677"

      def install
        bin.install "kit"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/kitproj/kit/releases/download/v0.1.19/kit_0.1.19_Darwin_x86_64.tar.gz"
      sha256 "8d1c07d5c1eebfdea3a581271c5fba5c864348b17d930949503a13869d1a959a"

      def install
        bin.install "kit"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/kitproj/kit/releases/download/v0.1.19/kit_0.1.19_Linux_arm64.tar.gz"
      sha256 "d5511627ce55ae166aef77180ef43b9cd470c08cb9ed3b01789b440be820ecd4"

      def install
        bin.install "kit"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/kitproj/kit/releases/download/v0.1.19/kit_0.1.19_Linux_x86_64.tar.gz"
      sha256 "8ad8d69f7dca5ffdee7a94993c872e682644d0af16274b2fd174e07d79d82cc9"

      def install
        bin.install "kit"
      end
    end
  end
end
