import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

import Swal from 'sweetalert2';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  bodyData: any;

  fullname: string = "";
  email: string = "";
  username: string = "";
  password: string = "";
  // role:string ="";


  constructor(private router: Router, private http: HttpClient) {
  }

  ngOnInit(): void {
  }




  register() {
    let bodyData =
    {

      // "img_user": this.img_user,
      "fullname": this.fullname,
      "email": this.email,
      "username": this.username,
      "password": this.password,
      // "role": this.role


    };



    if (!bodyData) {

      Swal.fire(
        'ไม่สำเร็จ!',
        'กรุณากรอกข้อมูลให้ครบถ้วน',
        'error'
      )
    }


    // if (this.password != this.confirmpassword) {
    //   Swal.fire(
    //     'ไม่สำเร็จ!',
    //     'กรุณากรอก password กับ confirmpassword ให้ตรงกัน',
    //     'warning'
    //   )

    // }




    else if (bodyData) {

      this.http.post("http://localhost:8880/register", bodyData).subscribe((resultData: any) => {


        console.log(resultData);
        Swal.fire(
          'Register สำเร็จ!',
          'ลงทะเบียนเรียบร้อยเเล้ว!',
          'success'

        )



        this.router.navigateByUrl('/login', {
          state: {

            // img_user: bodyData.img_user,
            fullname: bodyData.fullname,
            email: bodyData.email,
            username: bodyData.username,
            password: bodyData.password,
 
          }
        });
        // console.log(bodyData);
        // console.log("This is Firstname :" + bodyData.fullname);
        // console.log("This is Gmail :" + bodyData.email);
        // console.log("This is Password :" + bodyData.password);
        

      });


    }


  }

    save() {
      this.register();
    }


}