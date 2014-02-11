require 'spec_helper'

describe 'Home Page', type: :feature, js: true do

	before do
		visit('http://localhost:9000')
	end

	it 'shows welcome' do
		expect(page).to have_content('Welcome')
	end

end