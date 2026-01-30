export function getPreciseLocation(options = {}) {
  return new Promise((resolve, reject) => {
    if (!navigator.geolocation) {
      reject(new Error('geolocation_unavailable'));
      return;
    }
    const opts = {
      enableHighAccuracy: true,
      timeout: 10000,
      maximumAge: 0,
      ...options
    };
    navigator.geolocation.getCurrentPosition(
      pos => {
        const c = pos.coords;
        resolve({
          latitude: c.latitude,
          longitude: c.longitude,
          accuracy: c.accuracy,
          altitude: c.altitude,
          altitudeAccuracy: c.altitudeAccuracy,
          heading: c.heading,
          speed: c.speed,
          timestamp: pos.timestamp
        });
      },
      err => reject(err),
      opts
    );
  });
}

export function getCurrentLocation(options = {}) {
  return getPreciseLocation(options).then(c => ({
    latitude: c.latitude,
    longitude: c.longitude
  }));
}
