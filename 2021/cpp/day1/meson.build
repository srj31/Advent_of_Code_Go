app_file1 = files(
    'day1_A.cpp'
) 

conf_data = configuration_data()
# Set data
configure_file(input : 'input.txt',
  output : 'input.txt',
  configuration : conf_data)

configuration_inc = include_directories('.')

executable('day1_A', app_file1, include_directories: [incdir, configuration_inc])
