describe 'Dependencies' do
  describe 'Ruby', dependency: true, ruby: true do
    versions, latest = versions_for :ruby

    versions.each do |version|
      it "installs version #{version} correctly" do
        test_dep 'ruby', version: version, match: version
      end
    end

    it "installs version #{latest} as the default" do
      test_dep 'ruby', match: latest
    end
  end
end
