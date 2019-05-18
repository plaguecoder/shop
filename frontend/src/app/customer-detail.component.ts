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
  customer: {}

  constructor(private route: ActivatedRoute, private router: Router, private service: ConfigService) { }
  ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.id = params['id']
    })

    this.service.getCustomer(this.id)
      .subscribe((response) => {
        this.customer = response.body['data']
        this.transactions = this.customer['transactions']
        console.log('transa', this.transactions.length)
        let amount = 0
        for (let i = this.transactions.length-1; i != -1; i--) {
          console.log('i: ',i)
          
          if(this.transactions[i].type === 'CREDIT'){
            amount = amount + this.transactions[i].amount
            this.transactions[i]['total'] = amount
            //console.log(this.transactions[i]['total'])
          }
          if(this.transactions[i].type === 'DEBIT'){
            amount = amount - this.transactions[i].amount
            this.transactions[i]['total'] = amount
          }
        }
      })
  }

  goBack() {
    this.router.navigate(['/'])
  }

  goToAddTransactionPage() {
    this.router.navigate(['/customer-add-transaction', this.id])
  }
}
