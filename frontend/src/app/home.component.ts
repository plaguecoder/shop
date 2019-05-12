import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router } from '@angular/router'

@Component({
  selector: 'home',
  templateUrl: './home.component.html',
  providers: [ConfigService]
})
export class HomeComponent implements OnInit {

  customers = []
  breadcrumbList: Array<any> = [];

  constructor(private formBuilder: FormBuilder, private service: ConfigService, private router: Router) { }
  ngOnInit() {
    this.service.getCustomers().subscribe((response)=> {
    //  console.log(response)
      this.customers = response.body['data']
    })
  }

  goToAddCustomer() {
  //  console.log('inside')
    this.router.navigate(['/addCustomer'])
  }

  goToCustomerDetails(customerDetails) {
    this.router.navigate(['/customer-detail', customerDetails.id])
  }

}
