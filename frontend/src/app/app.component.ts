import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService} from './app.service'
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [ConfigService]
})
export class AppComponent implements OnInit {
  registerForm: FormGroup;
  submitted = false;


  constructor(private formBuilder: FormBuilder, private service: ConfigService) { }
  ngOnInit() {
    this.registerForm = this.formBuilder.group({
      name: ['', Validators.required],
      area: ['', Validators.required],
      amountDue: ['', [Validators.required,, Validators.minLength(1)]],
      phone: ['', [Validators.required, Validators.minLength(10)]]
    });
  }

  get f() { 
    //console.log(this.registerForm)
    return this.registerForm.controls; }

  onSubmit() {
    //console.log(this.f.phone.errors)
    this.submitted = true;

    // stop here if form is invalid
    if (this.registerForm.invalid) {
      return;
    }

    alert('SUCCESS!! :-)')
    this.service.addCustomer(this.registerForm.getRawValue())
    .subscribe((data) =>{
      console.log('response: ',data)
    })
  }
  submitData() {
    //console.log('customer: ', this.name, this.amountDue, this.phone, this.area)
  }

  // validateData() {
  //   if (this.name != '' && this.area != '' && this.phone.toString().length > 9 && this.phone.toString().length < 11 && this.amountDue != 0) {
  //     return false
  //   }
  //   return true
  // }
}
