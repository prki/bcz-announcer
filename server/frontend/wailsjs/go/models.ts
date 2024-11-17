export namespace main {
	
	export class CalloutUpdateMessage {
	    callout_id: string;
	    update_request: string;
	
	    static createFrom(source: any = {}) {
	        return new CalloutUpdateMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.callout_id = source["callout_id"];
	        this.update_request = source["update_request"];
	    }
	}
	export class SetupMatch {
	    game_name: string;
	    p1_name: string;
	    p2_name: string;
	    callout_id: string;
	
	    static createFrom(source: any = {}) {
	        return new SetupMatch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.game_name = source["game_name"];
	        this.p1_name = source["p1_name"];
	        this.p2_name = source["p2_name"];
	        this.callout_id = source["callout_id"];
	    }
	}

}

