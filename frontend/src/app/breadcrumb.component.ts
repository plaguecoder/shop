import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router, NavigationEnd, Route } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'breadcrumb',
    templateUrl: './breadcrumb.component.html',
    providers: [ConfigService]
})

export class BreadcrumbComponent implements OnInit {

    sub: any
    name: any
    bredcrumbs = {}
    bredcrumbList = []
    id: any

    constructor(private route: ActivatedRoute, private router: Router) {

        router.events.subscribe((event) => {
            if (event instanceof NavigationEnd) {
                this.name = event.url
                this.bredcrumbs[''] = "Home"
                this.bredcrumbs['addCustomer'] = "Add customer"
                this.bredcrumbs['customer-detail'] = "Customer details"
                this.bredcrumbs['customer-add-transaction'] = "Add transaction"
                this.bredcrumbs['areas'] = "Areas"
                this.bredcrumbs['add-area'] = "Add Area"
                let list = this.name.slice(1).split('/')
                if (list.length > 1) {
                    this.id = list[1]
                }
                for (let i in this.bredcrumbs) {
                    if (i == list[0]) {
                        let index = this.bredcrumbList.indexOf(this.bredcrumbs[i])

                        if (index === -1) {
                            this.bredcrumbList.push(this.bredcrumbs[i])
                        } else if (this.bredcrumbList.length > (index + 1)) {
                            this.bredcrumbList.splice(index + 1, this.bredcrumbList.length - (index + 1))
                        }
                    }
                }
            }
        })
    }

    ngOnInit() {

    }

    goToUrl(value) {
        console.log("value: ", value)
        for (let i in this.bredcrumbs) {
            if (value === this.bredcrumbs[i]) {
                console.log('i: ', '/' + i, this.id)
                if (i == "" || i=='areas' || i=='add-area') {
                    this.router.navigate(['/' + i])
                } else {
                    this.router.navigate(['/' + i, this.id])
                }

            }
        }

    }
}
