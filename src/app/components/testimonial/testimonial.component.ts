import { Component, OnInit, ViewChild } from '@angular/core';
import { SlickCarouselComponent } from 'ngx-slick-carousel';
import { MatDialog } from '@angular/material/dialog';
import { CartService} from '../../services/cart.service'
import { GetApiService } from './../../get-api.service';
import { MatIconModule } from '@angular/material/icon';
import { concatMap } from 'rxjs/operators';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-testimonial',
  templateUrl: './testimonial.component.html',
  styleUrls: ['./testimonial.component.css']
})
export class TestimonialComponent implements OnInit {
  @ViewChild('testimonialSlickModal') testimonialSlickModal: SlickCarouselComponent;

  testimonialData;

  constructor(
    public dialog: MatDialog,
    private cartService: CartService,
    private api:GetApiService
  ) { 
    
  }

  ngOnInit(): void {

    // this.testimonialData = (fruits as any).default;

     this.api.fetchUserTestimonialsApiCall().subscribe(
      (data1)=>{
        console.warn("get api data ", data1);
        this.testimonialData=data1;
      }
    )
  }

  testNext() {
    this.testimonialSlickModal.slickNext();
  }

  slideConfig = {
    "lazyLoad": 'progressive',
    'slidesToShow': 3,
    "slidesToScroll": 3,
    "arrows": true,
    "dots": true,
    speed: 300,
    "infinite": true,
    "responsive": [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 2,
          slidesToScroll: 2,
          infinite: true,
          dots: true
        }
      },
      {
        breakpoint: 760,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1
        }
      },
      {
        breakpoint: 480,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1
        }
      }
    ]
  };

  testPrev() {
    this.testimonialSlickModal.slickPrev();
  }
}
