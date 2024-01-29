import { Component, OnInit } from '@angular/core';
import AOS from 'aos';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit{


  ngOnInit() {

    //เพื่อไม่ให้มี error ของ aos (ใช้ npm i --save-dev @types/aos ในการติดตั้ง)
    if (typeof document !== 'undefined') {
      AOS.init();
    }
    
  }
}
