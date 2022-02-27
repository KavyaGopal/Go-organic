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
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
  fruitsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
  snacksApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
  vegetablesApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
  cosmeticsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
}
