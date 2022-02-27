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
    return this.httpClient.get('http://localhost:8000/getGroceries');
  }
  fruitsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getFruits');
  }
  snacksApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getSnacks');
  }
  vegetablesApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getVegetables');
  }
  cosmeticsApiCall()
  {
    return this.httpClient.get('http://localhost:8000/getCosmetics');
  }
}
