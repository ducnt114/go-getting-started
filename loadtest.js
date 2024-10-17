import http from 'k6/http';
import { check, sleep } from 'k6';

// Define options for the load test
export let options = {
//  vus: 5, // Virtual Users
//  duration: '20s', // Test duration
  stages: [
      { duration: '10s', target: 5 }, // ramp up to 400 users
      { duration: '30s', target: 100 }, // stay at 400 for ~4 hours
      { duration: '10s', target: 5 }, // scale down. (optional)
    ],
};

// Define the API endpoint you want to test
const url = 'http://localhost:8080/api/v1/user/1'; // Replace with your API URL

export default function () {
  // Perform a GET request
  let response = http.get(url);

  // Check if the response status is 200
  check(response, {
    'status is 200': (r) => r.status === 200,
    'response time is < 500ms': (r) => r.timings.duration < 500,
  });

  // Sleep for 1 second between iterations (optional)
//  sleep(1);
}