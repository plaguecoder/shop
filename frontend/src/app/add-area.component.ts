import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router, NavigationEnd, Route } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'add-area',
    templateUrl: './add-area.component.html',
    providers: [ConfigService]
})
export class AddAreaComponent implements OnInit {

    areaForm: FormGroup
    submitted = false
    constructor(private formBuilder: FormBuilder,private route: ActivatedRoute, private router: Router) {
       
    }

    ngOnInit() {
        this.areaForm = this.formBuilder.group({
            area: ['', Validators.required]
          });
    }

    get f() {
        //console.log(this.registerForm)
        return this.areaForm.controls;
      }
    
      onSubmit() {
        //console.log(this.f.phone.errors)
        this.submitted = true;
    
        // stop here if form is invalid
        if (this.areaForm.invalid) {
          return;
        }
        alert('Success!')
        // this.service.addCustomer(this.registerForm.getRawValue())
        //   .subscribe(response => {
        //     //console.log('response: ', response)
        //     if (response.status === 200) {
        //       this.submitted = false
        //       this.registerForm.reset()
        //       this.router.navigate(['/'])
        //     }
        //   })
      }
    
      goBack() {
        this.router.navigate(['/areas'])
      }

}
