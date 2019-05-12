import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'customer-add-transaction',
    templateUrl: './customer-add-transaction.component.html',
    providers: [ConfigService]
})
export class CustomerAddTransactionComponent implements OnInit {

    sub: any
    id: any
    transactions = []
    addTransaction: FormGroup;
    submitted = false;
    transactionTypes = ["DEBIT", "CREDIT"]

    constructor(private service: ConfigService, private route: ActivatedRoute, private router: Router, private formBuilder: FormBuilder) { }

    ngOnInit() {

        this.sub = this.route.params.subscribe(params => {
            this.id = params['id']
        })
        let today = new Date();
        let todayDate = today.getFullYear() + '-' + ('0' + (today.getMonth() + 1)).slice(-2) + '-' + ('0' + today.getDate()).slice(-2);
        this.addTransaction = this.formBuilder.group({
            date: [todayDate, Validators.required],
            type: ['', Validators.required],
            amount: [0, [Validators.required, Validators.minLength(1)]],
            description: [''],
            customer_id: [parseInt(this.id)]
        });
    }

    get f() {
        //console.log(this.registerForm)
        return this.addTransaction.controls;
    }

    get typeOfTransaction() {
        return this.addTransaction.get('type');
    }

    onSubmit() {
        this.submitted = true;
        //console.log(this.addTransaction.getRawValue(), this.f)
        // stop here if form is invalid
        if (this.addTransaction.invalid) {
            return;
        }
        this.service.addTransaction(this.addTransaction.getRawValue())
            .subscribe((response) => {
                if (response.status === 200) {
                    this.submitted = false
                    this.addTransaction.reset()
                    this.goBack()
                  }
            })
    }

    goBack() {
        this.router.navigate(['/customer-detail', this.id])
    }
}
