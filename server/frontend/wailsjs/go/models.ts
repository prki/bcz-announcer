export namespace main {
	
	export class SetupMatch {
	    p1_name: string;
	    p2_name: string;
	    callout_id: string;
	
	    static createFrom(source: any = {}) {
	        return new SetupMatch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.p1_name = source["p1_name"];
	        this.p2_name = source["p2_name"];
	        this.callout_id = source["callout_id"];
	    }
	}

}

