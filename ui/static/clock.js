function updateClock() {
  const now = new Date();

  const options = {
    //timeZone: 'America/New_York',
    hour: 'numeric',
    minute: '2-digit',
    second: '2-digit',
    hour12: true,
  }

  const timeString = now.toLocaleTimeString('en-US', options);
  const timezone = now.toLocaleTimeString('en-US', {
    timeZone: 'America/New_York',
    timeZoneName: 'short'
  }).split(' ').pop();

  document.getElementById('current-time').textContent = `${timeString}`;
}

updateClock();

setInterval(updateClock, 1000); 
  
