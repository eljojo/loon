describe 'Dependencies' do
  describe 'Golang', dependency: true, golang: true do
    versions, latest = versions_for :golang

    versions.each do |version|
      it "installs version #{version} correctly" do
        test_dep 'golang', version: version, cmd: 'go version', match: version
      end
    end

    it "installs version #{latest} as the default" do
      test_dep 'golang', cmd: 'go version', match: latest
    end
  end
end
