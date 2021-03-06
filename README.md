retail-4-5-4-calendar
=====================

Program that generates the retail 4-5-4 calendar dates that correspond to the Gregorian calendar.  
This program starts at January 30, 2000 (Gregorian) because that is Day 1 Week 1 of the 2000 Retail 4-5-4 calendar. 
Due to the nature of the pattern that the 4-5-4 calendar is generated by, it must start generating from 
Jan 30th 2000. 

Info on retail 4-5-4 calendar at:  
[National Retail Federation](http://www.nrf.com/modules.php?name=Pages&sp_id=391)  

Currently the NRF only supplies the 4-5-4 calendar approximately 4 years in advance. This program will generate dates 
however far into the future you would like.

This will output data to "output.txt" in the same directory that you run the program from. 
Carriage returns don't display properly in notepad, so use notepad++ or something similar 
to view output.

**Use**:  
Specify the amount of years to output with the years variable  
Currently the output is sent to "output.txt"  
The output shows the Gregorian year, month, and date on the far left 
and the Retail 4-5-4 year, week of year, and day of week in the next
three columns

**To Do**:  
- Enhance output format  
- Allow user to specify a range of dates to be printed  
- Implement some sort of GUI
- Support Holidays and Sales Release Dates
