import { Component } from '@angular/core';
import { PostService } from './post.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'admin-panel-layout';
  sideBarOpen = true;
  constructor(private post:PostService){

  }
  

  sideBarToggler() {
    this.sideBarOpen = !this.sideBarOpen;
  }
  p:any;
  data:any=[];
  getData(){
    this.post.getData().subscribe(
      (data) =>{
        this.data=data;
     
      }
    );
  }
}