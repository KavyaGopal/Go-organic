import { Component } from '@angular/core';
import { GetApiService } from './get-api.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Go-organic';
  
  constructor(
    private api:GetApiService
  ){

  }

  //this calls api on page load
  ngOnInit(){
    this.api.apiCall().subscribe(
      (data)=>{
        console.warn("get api data ", data);
      }
    )
  }
}
