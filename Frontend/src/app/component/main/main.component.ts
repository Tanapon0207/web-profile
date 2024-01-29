import { Component } from '@angular/core';
// import AOS from 'aos';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent {

  isMenuActive: boolean = false;

  toggleHam(): void {
    this.isMenuActive = !this.isMenuActive;
  }
  
}
