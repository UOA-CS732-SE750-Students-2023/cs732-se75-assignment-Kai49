export function getMonthDates(date) {


  const year = date.getUTCFullYear();


  const month = date.getMonth() ;
  

  
  const firstDay = new Date(Date.UTC(year, month, 1));
  const lastDay = new Date(Date.UTC(year, month + 1, 0));
  const daysInMonth = lastDay.getUTCDate();
  const dates = [];

  // Find the first day of the week for the current month
  const firstDayOfWeek = firstDay.getUTCDay();

  // Add empty days before the first day of the month to align the calendar correctly
  for (let i = 0; i < firstDayOfWeek; i++) {
    dates.push({
      day: null,
      date: null,
    });
  }

  // Add actual days of the month
  for (let i = 1; i <= daysInMonth; i++) {
    const dayDate = new Date(Date.UTC(year, month, i));
    dates.push({
      day: dayDate.getUTCDate(),
      date: dayDate.toISOString().split('T')[0],
    });
  }
   

  return dates;
  }
  