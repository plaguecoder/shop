import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router, NavigationEnd, Route } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'areas',
    templateUrl: './areas.component.html',
    providers: [ConfigService]
})
export class AreasComponent implements OnInit {
    areas = []
    constructor(private route: ActivatedRoute, private router: Router) {
        this.areas.push('Murgeshpalya')
        this.areas.push('Kodihalli')
        this.areas.push('Domlur')
        this.areas.push("kormangala")
        this.areas.push('BTM')
        this.areas.push('JP nagar')
    }

    ngOnInit() {

    }

    goToAddArea() {
        this.router.navigate(['/add-area'])
    }
}
