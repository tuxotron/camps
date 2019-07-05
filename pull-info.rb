#!/usr/bin/env ruby

require 'nokogiri'
require 'open-uri'
require 'json'

class Camp
  attr_accessor :name, :street1, :street2, :city, :state, :zip, :phone, :fax

  def as_json(options = {})
    {
        name: @name,
        street1: @street1,
        street2: @street2,
        city: @city,
        state: @state,
        zip: @zip,
        phone: @phone,
        fax: @fax
    }
  end

  def to_json(*options)
    as_json(*options).to_json(*options)
  end
end

doc = Nokogiri::HTML(open('http://concentrationcamps.us/cbp.html'))

camp = nil
camps = Array.new

doc.css('body > p.Default_20_Text > span').map {|node|
  if node['class'] == 'T1'
    if camp != nil
      camps.push(camp)
    end

    camp = Camp.new
    camp.name = node.text
  elsif node['class'] == 'T2'
    if node.text.start_with?('Phone')
      camp.phone = node.text.gsub('Phone: ', '')
    elsif node.text.start_with?('Fax')
      camp.fax = node.text.gsub('Fax: ', '')
    elsif node.text.include?(',')
      addr = node.text.split(',')
      camp.city = addr[0]
      camp.state = addr[1].split[0]
      camp.zip = addr[1].split[1]
    else
      if camp.street1.to_s.strip.empty?
        camp.street1 = node.text
      else
        camp.street2 = node.text
      end
    end
  end
}

puts "{ \"camps\": #{camps.to_json} }"