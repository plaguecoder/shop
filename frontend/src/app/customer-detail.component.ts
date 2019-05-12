import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
  selector: 'customer-detail',
  templateUrl: './customer-detail.component.html',
  providers: [ConfigService]
})
export class CustomerDetailComponent implements OnInit {

  sub: any
  id: any
  transactions = []
  customer : {}

  constructor(private route: ActivatedRoute, private router: Router, private service: ConfigService) {}
  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
        this.id = params['id']
    })

    this.service.getCustomer(this.id)
        .subscribe((response) => {
         //   console.log('customer', response)
            this.customer = response.body['data']
            this.transactions = this.customer['transactions']
        })
  }

  goBack() {
    this.router.navigate(['/'])
  }

  goToAddTransactionPage(){
      this.router.navigate(['/customer-add-transaction', this.id])
  }
}
