# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kit < Formula
  desc "Kitful local dev."
  homepage "https://github.com/alexec/kit"
  version "0.0.30"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/alexec/kit/releases/download/v0.0.30/kit_0.0.30_Darwin_x86_64.tar.gz"
      sha256 "b207906808c4d30aa9162f8cb925335a1a6d02e89e0578cbbf51775e8a581713"

      def install
        bin.install "kit"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/alexec/kit/releases/download/v0.0.30/kit_0.0.30_Darwin_arm64.tar.gz"
      sha256 "f0809c32a6afacd18c8d01fbfec4617f828660a86ce9178399d185aed1fefe93"

      def install
        bin.install "kit"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/alexec/kit/releases/download/v0.0.30/kit_0.0.30_Linux_arm64.tar.gz"
      sha256 "3268fb0b2ee7ff43245b448bb495e82eb30cb88a30a4889983e0a5cdc78270e6"

      def install
        bin.install "kit"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/alexec/kit/releases/download/v0.0.30/kit_0.0.30_Linux_x86_64.tar.gz"
      sha256 "7a00f7732f4abea732a0e7f7c93b39be9a7b4dc09ea18a4ac75caa44812cfc52"

      def install
        bin.install "kit"
      end
    end
  end
end
