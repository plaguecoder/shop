import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router } from '@angular/router'
@Component({
  selector: 'app-root',
  templateUrl: './add-customer.component.html',
  styleUrls: ['./add-customer.component.css'],
  providers: [ConfigService]
})
export class AddCustomerComponent implements OnInit {
  registerForm: FormGroup;
  submitted = false;
  success = false;


  constructor(private formBuilder: FormBuilder, private service: ConfigService, private router: Router) { }
  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      name: ['', Validators.required],
      area: ['', Validators.required],
      amountDue: ['', [Validators.required, , Validators.minLength(1)]],
      phone: ['', [Validators.required, Validators.minLength(10)]],
      description: ['']
    });
  }

  get f() {
    //console.log(this.registerForm)
    return this.registerForm.controls;
  }

  onSubmit() {
    //console.log(this.f.phone.errors)
    this.submitted = true;

    // stop here if form is invalid
    if (this.registerForm.invalid) {
      return;
    }

    this.service.addCustomer(this.registerForm.getRawValue())
      .subscribe(response => {
        //console.log('response: ', response)
        if (response.status === 200) {
          this.submitted = false
          this.registerForm.reset()
          this.router.navigate(['/'])
        }
      })
  }

  goBack() {
    this.router.navigate(['/'])
  }
}
