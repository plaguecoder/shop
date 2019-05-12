import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService} from './app.service'
import {Router} from '@angular/router'

@Component({
  selector: 'home',
  templateUrl: './home.component.html',
  providers: [ConfigService]
})
export class HomeComponent implements OnInit {

  constructor(private formBuilder: FormBuilder, private service: ConfigService, private router: Router) { }
  ngOnInit() {
  }

  goToAddCustomer(){
    console.log('inside')
    this.router.navigate(['/addCustomer'])
  }
}
