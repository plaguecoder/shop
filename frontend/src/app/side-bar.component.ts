import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConfigService } from './app.service'
import { Router, NavigationEnd, Route } from '@angular/router'
import { ActivatedRoute } from '@angular/router'

@Component({
    selector: 'sidebar',
    templateUrl: './side-bar.component.html',
    providers: [ConfigService]
})
export class SidebarComponent implements OnInit {

    constructor(private route: ActivatedRoute, private router: Router) {

    }

    ngOnInit() {

    }

}
