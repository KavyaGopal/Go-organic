import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class GetApiService {

  constructor(
    private httpClient: HttpClient
  ) { }

  //call the GET api backend
  // apiCall()
  // {
  //   return this.httpClient.get('http://localhost:8000/getFruits');
  // }
  groceriesApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchProduct/Groceries');
  }
  fruitsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchProduct/Fruits');
  }
  snacksApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchProduct/Snacks');
  }
  vegetablesApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchProduct/Vegetables');
  }
  cosmeticsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchProduct/Cosmetics');
  }

  fetchUserTestimonialsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/api/fetchUserTestimonials');
  }
}
